import { expect, test } from './fixtures/auth.fixture';

test.describe('audit e2e', () => {
  test('日志页面加载并可执行筛选', async ({ loginAsAdmin, page }) => {
    await loginAsAdmin();
    await page.goto('/#/audit/baselogs', { waitUntil: 'domcontentloaded' });

    await expect(page.getByText('ID').first()).toBeVisible();
    await page.getByText(/Search|搜索/).first().click();

    await page.locator('input').first().fill('e2e-conn');
    await page.locator('button').filter({ hasText: /Search|搜索/ }).last().click();

    await expect(page.getByText('ID').first()).toBeVisible();
  });
});
