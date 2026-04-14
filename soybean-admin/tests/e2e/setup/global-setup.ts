import process from 'node:process';
import type { FullConfig } from '@playwright/test';

async function checkResponseOk(url: string, init?: RequestInit) {
  const resp = await fetch(url, init);
  if (!resp.ok) {
    throw new Error(`request failed: ${url}, status=${resp.status}`);
  }
  return resp;
}

async function globalSetup(_config: FullConfig) {
  const backendBase = process.env.E2E_BACKEND_URL || 'http://127.0.0.1:12345';
  const username = process.env.E2E_ADMIN_USER || 'admin';
  const password = process.env.E2E_ADMIN_PASS || 'admin123456';

  await checkResponseOk(`${backendBase}/admin/auth/captcha`);

  const loginResp = await checkResponseOk(`${backendBase}/admin/auth/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      username,
      password,
      code: '0000',
      captchaId: 'e2e-bypass'
    })
  });

  const loginJson = (await loginResp.json()) as { code?: number; data?: { token?: string } };
  if (loginJson.code !== 200 || !loginJson.data?.token) {
    throw new Error('admin login check failed in global setup');
  }
}

export default globalSetup;
