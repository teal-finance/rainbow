import { test, expect } from '@playwright/test';

test('SOL preset test', async ({ page }) => {
  let setPageReady: (v: unknown) => void;
  const onPageReady = new Promise((resolve) => setPageReady = resolve)
  page.on("requestfinished", async (request) => {
    if (request.url() == "http://localhost:8090/graphql") {
      setPageReady(true)
    }
  })
  await page.goto('/');
  await onPageReady;
  const btc = page.locator('#BTC input[type=checkbox]');
  const sol = page.locator('#SOL input[type=checkbox]');
  expect(await btc.isChecked()).toBeTruthy();
  expect(await sol.isChecked()).toBeTruthy();
  await page.click("#presets-select")
  await page.click("#preset-SOL")
  expect(await btc.isChecked()).toBeFalsy();
  expect(await sol.isChecked()).toBeTruthy();
  await page.click("#presets-select")
  await page.click("#preset-All")
  expect(await btc.isChecked()).toBeTruthy();
  expect(await sol.isChecked()).toBeTruthy();
});