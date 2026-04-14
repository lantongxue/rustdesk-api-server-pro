import process from 'node:process';
import { test as base, expect } from '@playwright/test';

type AuthFixtures = {
  loginAsAdmin: () => Promise<void>;
};

export const test = base.extend<AuthFixtures>({
  loginAsAdmin: async ({ page, baseURL }, use) => {
    await use(async () => {
      const username = process.env.E2E_ADMIN_USER || 'admin';
      const password = process.env.E2E_ADMIN_PASS || 'admin123456';
      const loginPath = `${baseURL}/login/pwd-login`;

      await page.goto(loginPath, { waitUntil: 'domcontentloaded' });
      await page.getByPlaceholder(/Please enter user name|请输入用户名/).fill(username);
      await page.getByPlaceholder(/Please enter password|请输入密码/).fill(password);
      // Captcha is bypassed when backend runs with E2E_SKIP_CAPTCHA=true.
      await page.getByPlaceholder(/Please enter verification code|请输入验证码/).fill('0000');
      await page.getByRole('button', { name: /Confirm|确定|确认/ }).click();
      await page.waitForURL(/\/home/, { timeout: 20_000 });
      await expect(page.getByText('Rustdesk Api Server')).toBeVisible();
    });
  }
});

export { expect };
