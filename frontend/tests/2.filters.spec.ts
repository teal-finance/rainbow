import { test, expect } from '@playwright/test';

test('BTC filter test', async ({ page }) => {
  let setPageReady: (v: unknown) => void;
  const onPageReady = new Promise((resolve) => setPageReady = resolve)
  page.on("requestfinished", async (request) => {
    if (request.url() == "http://localhost:8090/graphql") {
      setPageReady(true)
    }
  })
  await page.goto('/');
  await onPageReady;
  const i = page.locator('#BTC input[type=checkbox]');
  expect(await i.isChecked()).toBeTruthy()
  await page.click('#BTC');
  expect(await i.isChecked()).toBeFalsy()
});