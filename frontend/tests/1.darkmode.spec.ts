import { test, expect } from '@playwright/test';

test('dark mode switch test', async ({ page }) => {
  await page.goto('/');
  const el = page.locator('#toplevel');
  const body = page.locator('body');
  expect(await el.getAttribute("class")).not.toContain("dark")
  await page.click('#dark-mode-switch');
  expect(await el.getAttribute("class")).toContain("dark")
  expect(await body.getAttribute("class")).toBe("bg-background-dark text-foreground-dark")
  await page.click('#dark-mode-switch');
  expect(await el.getAttribute("class")).not.toContain("dark")
});
