import { gfm } from 'turndown-plugin-gfm';
import TurndownService from 'turndown';

// TODO: replace with JSDom once JSDom works with ESModules https://github.com/jsdom/jsdom/issues/2475
// OR, find out why h2 element can't be selected properly
import { Window } from 'happy-dom';

export const capturePage = (htmlString: string, originalUrl: string) => {
  const window = new Window();
  const originalDocument = window.document;

  const newWindow = new Window();
  const newDocument = newWindow.document;

  originalDocument.documentElement.replaceWith(htmlString);

  const articles = originalDocument.querySelectorAll('article');
  for (const [i, article] of articles.entries()) {
    // I have no idea why article.querySelector and getElementsByTagName don't work here...
    // const headerElement = article.querySelector('h2');
    const headerElement = article.children.filter(
      (c) => c.nodeName === 'H2'
    )[0];
    const headerText = headerElement.textContent.replaceAll('---', '').trim();

    const partHeader = newDocument.createElement('h3');
    if (i === 0) {
      const originalPageLink = newDocument.createElement('a');
      originalPageLink.setAttribute('href', originalUrl);
      originalPageLink.textContent = headerText;

      const promptHeader = newDocument.createElement('h1');
      promptHeader.innerHTML = originalPageLink.outerHTML;
      newDocument.body.appendChild(promptHeader);

      const descriptionHeader = newDocument.createElement('h2');
      descriptionHeader.textContent = 'Description';
      newDocument.body.appendChild(descriptionHeader);

      partHeader.textContent = 'Part One';
    } else {
      partHeader.textContent = headerText;
    }

    newDocument.body.appendChild(partHeader);

    article.removeChild(headerElement);

    newDocument.body.appendChild(article);

    article.querySelectorAll('code > em:only-child').map((e) => {
      const newCodeElem = originalDocument.createElement('code');
      newCodeElem.innerHTML = e.textContent;

      e.parentNode.parentNode.replaceChild(newCodeElem, e.parentNode);
    });
  }

  const turndownService = new TurndownService({
    headingStyle: 'atx'
  });
  turndownService.use(gfm);
  turndownService.keep(['span']);

  return turndownService.turndown(newDocument.body.innerHTML).concat('\n');
};
