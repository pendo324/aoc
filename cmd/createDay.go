package cmd

import (
	"bytes"
	_ "embed"
	"errors"
	"text/template"

	"context"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/tools/go/ast/astutil"
)

type ctxKey string

func NewCreateDayCmd(rootCmd *cobra.Command) {
	createDayCmd := &cobra.Command{
		Use:  "create-day",
		RunE: createDay,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
				return err
			}
			if err := cobra.MaximumNArgs(3)(cmd, args); err != nil {
				return err
			}
			if len(args[0]) != 4 {
				return fmt.Errorf("invalid year: %s", args[0])
			}
			year, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid year: %s", args[0])
			}
			if len(args[1]) != 1 && len(args[1]) != 2 {
				return fmt.Errorf("invalid day: %s", args[1])
			}
			day, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid day: %s", args[1])
			}

			ctx := cmd.Context()
			if len(args) == 3 {
				ctx = context.WithValue(ctx, ctxKey("basePath"), args[2])
			} else {
				ex, err := os.Executable()
				if err != nil {
					return fmt.Errorf("failed to get current executable: %w", err)
				}
				ctx = context.WithValue(ctx, ctxKey("basePath"), filepath.Dir(ex))
			}

			ctx = context.WithValue(ctx, ctxKey("year"), year)
			ctx = context.WithValue(ctx, ctxKey("day"), day)
			cmd.SetContext(ctx)

			return nil
		},
	}

	rootCmd.AddCommand(createDayCmd)
}

func createDay(cmd *cobra.Command, _ []string) error {
	year := cmd.Context().Value(ctxKey("year")).(int)
	day := cmd.Context().Value(ctxKey("day")).(int)
	basePath := cmd.Context().Value(ctxKey("basePath")).(string)

	sessionCookie := viper.GetString("aoc_session_cookie")
	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	}

	client := &http.Client{}
	baseUrl := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	inputUrl := fmt.Sprintf("%s/input", baseUrl)

	baseReq, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	inputReq, err := http.NewRequest("GET", inputUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	baseReq.AddCookie(cookie)
	inputReq.AddCookie(cookie)

	baseRes, err := client.Do(baseReq)
	if err != nil {
		return fmt.Errorf("failed to get baseUrl (%s): %w", baseUrl, err)
	}

	inputRes, err := client.Do(inputReq)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}

	readme, err := convertToMarkdown(baseRes.Body, baseUrl)
	if err != nil {
		return fmt.Errorf("failed to convert to markdown: %w", err)
	}

	if err := baseRes.Body.Close(); err != nil {
		return fmt.Errorf("failed to close baseRes body: %w", err)
	}

	if err := write(inputRes.Body, readme, basePath, year, day); err != nil {
		return fmt.Errorf("failed to write input: %w", err)
	}

	if err := inputRes.Body.Close(); err != nil {
		return fmt.Errorf("failed to close inputRes body: %w", err)
	}

	if err := updateImport(year, day); err != nil {
		return fmt.Errorf("failed to update imports: %w", err)
	}

	return nil
}

//go:embed day.go.TEMPLATE
var dayTmpl string

type DayFields struct {
	Day     int
	FullDay string
	Year    int
}

func write(input io.ReadCloser, readme, basePath string, year, day int) error {
	yearDayPath := filepath.Join(basePath, "years", strconv.Itoa(year), fmt.Sprintf("%02d", day))

	if err := os.MkdirAll(yearDayPath, 0o644); err != nil {
		return fmt.Errorf("failed to create year/day directory: %w", err)
	}

	if err := os.WriteFile(filepath.Join(yearDayPath, "README.md"), []byte(readme), 0o644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	tmpl, err := template.New("dayTmpl").Parse(dayTmpl)
	if err != nil {
		return fmt.Errorf("failed to create template: %w", err)
	}
	var dayGo bytes.Buffer
	if err = tmpl.Execute(&dayGo, DayFields{
		Day:     day,
		FullDay: fmt.Sprintf("%02d", day),
		Year:    year,
	}); err != nil {
		return fmt.Errorf("failed to execute template file: %w", err)
	}

	// Only write main template if it doesn't already exist
	mainPath := filepath.Join(yearDayPath, "main.go")
	if _, err := os.Stat(mainPath); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(mainPath, dayGo.Bytes(), 0o644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	inputFile, err := os.Create(filepath.Join(yearDayPath, "input"))
	if err != nil {
		return fmt.Errorf("failed to create input file: %w", err)
	}
	b, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}
	if _, err := inputFile.WriteString(strings.TrimRight(string(b), "\n")); err != nil {
		return fmt.Errorf("failed to write to input file: %w", err)
	}
	if err := inputFile.Close(); err != nil {
		return fmt.Errorf("failed to close input file: %w", err)
	}

	return nil
}

func convertToMarkdown(body io.ReadCloser, url string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", fmt.Errorf("failed to create doc: %w", err)
	}

	out := ""
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		headerText := strings.TrimSpace(strings.ReplaceAll(s.Find("h2").Text(), "---", ""))
		if i == 0 {
			out = fmt.Sprintf("# [%s](%s)\n\n", headerText, url)
			out = fmt.Sprintf("%s## Description\n\n", out)
			out = fmt.Sprintf("%s### Part One\n\n", out)
		} else {
			out = fmt.Sprintf("%s### %s\n\n", out, headerText)
		}

		// remove header since it has been added to document manually already
		s.Find("h2").Remove()

		// remove extra new line from code blocks
		s.Find("pre > code").Each(func(_ int, s *goquery.Selection) {
			s.SetText(strings.TrimRight(s.Text(), "\n"))
		})

		converter := md.NewConverter("", true, nil)
		out = fmt.Sprintf("%s%s\n\n", out, converter.Convert(s))
	})

	return out, nil
}

func updateImport(year, day int) error {
	baseImportPath := "github.com/pendo324/aoc/years"
	importPath := fmt.Sprintf("%s/%d/%02d", baseImportPath, year, day)

	yearsGoPath := filepath.Join(".", "years", "years.go")

	b, err := os.ReadFile(yearsGoPath)
	if err != nil {
		return fmt.Errorf("failed to read years file: %w", err)
	}

	fset := token.NewFileSet() // positions are relative to fset
	af, err := parser.ParseFile(fset, "", b, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to parse file: %w", err)
	}

	astutil.AddNamedImport(fset, af, "_", importPath)
	// ast.SortImports(fset, af)

	if err := os.Truncate(yearsGoPath, 0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	var out bytes.Buffer
	if err := format.Node(&out, fset, af); err != nil {
		return err
	}

	if err := os.WriteFile(yearsGoPath, out.Bytes(), 0o644); err != nil {
		return err
	}

	return nil
}
