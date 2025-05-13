import clsx from "clsx"

type Props = Readonly<{
  text: string | null | undefined
  kind: 'success' | 'error'
}>

export function Message({ text, kind }: Props) {
  if (!text) return null
  return (
    <div className="fixed top-2 left-0 right-0 flex justify-center items-center">
      <div
        role="alert"
        className={clsx(
          "p-2 rounded-lg text-white",
          {
            "bg-green-500": kind === "success",
            "bg-red-500": kind === "error"
          }
        )}
      >
        {text}
      </div>
    </div>
  )
}
