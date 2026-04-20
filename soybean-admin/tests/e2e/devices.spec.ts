import { expect, test } from './fixtures/auth.fixture';

test.describe('devices e2e', () => {
  test('设备列表加载与筛选', async ({ loginAsAdmin, page }) => {
    await loginAsAdmin();
    await page.goto('/#/devices', { waitUntil: 'domcontentloaded' });

    await expect(page.getByText('ID').first()).toBeVisible();
    await page.getByText(/Search|搜索/).first().click();

    const searchInput = page.locator('input').first();
    await searchInput.fill('e2e-host');
    await page.locator('button').filter({ hasText: /Search|搜索/ }).last().click();

    await expect(page.getByText('ID').first()).toBeVisible();
  });
});
