import clsx from "clsx";
import { PropsWithChildren } from "react";

type Props = Readonly<PropsWithChildren<{
  errors?: string[]
  required?: boolean
  id: string
  label: string
}>>

export function Field({ children, id, label, required, errors = [] }: Props) {
  return (
    <fieldset className="flex flex-col gap-1">
      <label htmlFor={id}>{label} {required && <span className="text-red-400">*</span>}</label>
      <div className={clsx("border rounded-md p-2", { "border-red-500 text-red-500": errors.length, "border-foreground": !errors.length })}>
        {children}
      </div>
      {errors.map((message, index) => (
        <span key={index} className="text-red-500">{message}</span>
      ))}
    </fieldset>
  )
}
