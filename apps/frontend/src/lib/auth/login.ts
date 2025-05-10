import { ApiResponseError, ValidationError } from "../error";
import { err, ok } from "../result";
import { AuthLoginPayload, AuthLoginResult } from "./schema";

export async function login(form: FormData) {

  const { success, data, error } = AuthLoginPayload.safeParse({
    username: form.get("username"),
    password: form.get("password")
  })

  if (!success) {
    return err(ValidationError.fromError(error))
  }

  const response = await fetch("http://localhost:4000/auth", {
    method: "post",
    headers: new Headers({
      "content-type": "application/x-www-form-urlencoded",
      "accept": "application/json",
    }),
    body: new URLSearchParams(Object.entries(data)).toString()
  })

  if (!response.ok) {
    return err(await ApiResponseError.fromResponse(response))
  }

  return ok<AuthLoginResult>(await response.json())
}
