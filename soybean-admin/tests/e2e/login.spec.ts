import process from 'node:process';
import { expect, test } from './fixtures/auth.fixture';

test.describe('login e2e', () => {
  test('登录成功并跳转首页', async ({ loginAsAdmin, page }) => {
    await loginAsAdmin();
    await expect(page).toHaveURL(/\/home/);
  });

  test('错误密码登录失败', async ({ page, baseURL }) => {
    const loginPath = `${baseURL}/login/pwd-login`;
    await page.goto(loginPath, { waitUntil: 'domcontentloaded' });

    await page.getByPlaceholder(/Please enter user name|请输入用户名/).fill(process.env.E2E_ADMIN_USER || 'admin');
    await page.getByPlaceholder(/Please enter password|请输入密码/).fill('wrong-password');
    await page.getByPlaceholder(/Please enter verification code|请输入验证码/).fill('0000');
    await page.getByRole('button', { name: /Confirm|确定|确认/ }).click();

    await expect(page).toHaveURL(/\/login/);
  });
});
