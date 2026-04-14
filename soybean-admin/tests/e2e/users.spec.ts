import { expect, test } from './fixtures/auth.fixture';

test.describe('users e2e', () => {
  test('用户列表加载并新增用户', async ({ loginAsAdmin, page }) => {
    await loginAsAdmin();
    await page.goto('/#/user/list', { waitUntil: 'domcontentloaded' });

    await expect(page.getByText('ID').first()).toBeVisible();

    const suffix = Date.now().toString().slice(-6);
    const username = `e2e_user_${suffix}`;
    const password = `E2ePass_${suffix}`;
    const nickname = `E2E-${suffix}`;
    const email = `e2e_${suffix}@example.com`;

    await page.getByRole('button', { name: /Add|添加|新增/ }).click();
    const drawer = page.locator('.n-drawer').last();
    await drawer.getByPlaceholder(/Input Username|请输入用户名/).fill(username);
    await drawer.getByPlaceholder(/Input Password|请输入密码/).fill(password);
    await drawer.getByPlaceholder(/Input Nickname|请输入昵称/).fill(nickname);
    await drawer.locator('.n-form-item').nth(3).locator('input').first().fill(email);
    await drawer.getByRole('button', { name: /Confirm|确定|确认/ }).click();

    await expect(page.getByText(username)).toBeVisible({ timeout: 20_000 });
  });
});
