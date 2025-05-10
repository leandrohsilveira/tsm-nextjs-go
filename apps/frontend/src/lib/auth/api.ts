import { api, parsers } from "../api";
import type { AuthLoginPayload, AuthLoginResult } from "./schema";

export async function login(data: AuthLoginPayload) {
  return await api.post('/auth', parsers.request.json(data), parsers.response.json<AuthLoginResult>())

}
