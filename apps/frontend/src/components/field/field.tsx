import clsx from 'clsx'
import { PropsWithChildren } from 'react'
import styles from './field.module.css'

type Props = Readonly<
  PropsWithChildren<{
    errors?: string[]
    id: string
    label: string
  }>
>

export function Field({ children, id, label, errors = [] }: Props) {
  return (
    <fieldset className={clsx('flex flex-col gap-1', styles.field)}>
      <label htmlFor={id}>{label}</label>
      <div
        className={clsx('border rounded-md p-2', styles.field_target, {
          'border-red-500 text-red-500': errors.length,
          'border-foreground': !errors.length,
        })}
      >
        {children}
      </div>
      {errors.map((message, index) => (
        <span key={index} className="text-red-500">
          {message}
        </span>
      ))}
    </fieldset>
  )
}
