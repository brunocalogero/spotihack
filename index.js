const puppeteer = require("puppeteer");
const { TimeoutError } = require("puppeteer/Errors");

const SONG = "getting it back";
const QUERY_SELECTOR = "#query";
const FORM_SELECTOR = "#sort";
const SEARCH_LIST_SELECTOR = "#result > div.list-group";
const DOWNLOAD_SELECTOR =
  "body > div.wrapper > div.container > div > span > button";

puppeteer.launch().then(async browser => {
  // let err, nav;
  const page = await browser.newPage();
  await to(page.goto("https://myfreemp3c.com/music"));

  await to(page.click(QUERY_SELECTOR));
  await to(page.keyboard.type(SONG));

  // await page.waitForNavigation();
  await page.screenshot({ path: "screenshot.png" });

  await page.click(DOWNLOAD_SELECTOR);
  try {
    await page.waitFor(2 * 1000); // need to find better way
    // await page.waitForSelector("SEARCH_LIST_SELECTOR");
    // await page.waitForNavigation(0, "domcontentloaded");
  } catch (e) {
    if (e instanceof TimeoutError) {
      throw e;
    }
  }

  await page.screenshot({ path: "screenshot2.png" });
  // await page.screenshot({ path: "screenshot2.png" });
  // await page.screenshot({ path: "screenshot2.png" });
  // await page.waitForNavigation();

  await browser.close();
});

function to(promise) {
  return promise
    .then(data => {
      return [null, data];
    })
    .catch(err => [err]);
}
