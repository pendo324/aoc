import cheerio from 'cheerio';
import { gfm } from 'turndown-plugin-gfm';
import TurndownService from 'turndown';

const { load } = cheerio;

export const capturePage = (body: string, originalUrl: string) => {
  const document = load(body);
  const newDocument = load('');

  let dayTitle: string;
  document(`article`).each((index, article) => {
    const headingElement = document(article).find('h2');
    let headingElementText = headingElement.text();
    let newHeadingElementText = headingElementText.replaceAll('---', '').trim();

    if (index === 0) {
      dayTitle = newHeadingElementText;
      newHeadingElementText = 'Part One';
    }

    headingElement.replaceWith(`<h3>${newHeadingElementText}</h3>`);

    document(article)
      .find('code > em:only-child')
      .each((_, emInCodeBlock) => {
        document(emInCodeBlock.parentNode).replaceWith(
          `<code>${document(emInCodeBlock).text()}</code>`
        );
      });
  });

  newDocument
    .root()
    .append(`<h1>${dayTitle}</h1>`)
    .append('<h2>Description</h2>')
    .append(document('article'));

  const turndownService = new TurndownService({
    headingStyle: 'atx',
  });
  turndownService.use(gfm);
  turndownService.keep(['span']);

  return turndownService.turndown(newDocument.html()).concat('\n');
};
