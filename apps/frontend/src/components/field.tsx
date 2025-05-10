import { PropsWithChildren } from "react";

type Props = Readonly<PropsWithChildren<{
  error?: string
  required?: boolean
  id: string
  label: string
}>>

export function Field({ children, id, label, required }: Props) {
  return (
    <fieldset className="flex flex-col gap-1">
      <label htmlFor={id}>{label} {required && <span className="text-red-400">*</span>}</label>
      <div className="border border-foreground rounded-md p-2">
        {children}
      </div>
    </fieldset>
  )
}
