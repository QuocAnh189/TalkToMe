import { InputHTMLAttributes, forwardRef } from 'react'
import { cn } from '@utils/cn'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  error?: string
  label?: string
  helperText?: string
}

const Input = forwardRef<HTMLInputElement, InputProps>(({ className, error, label, helperText, ...props }, ref) => {
  return (
    <div className="form-control w-full">
      {label && (
        <label className="label">
          <span className="label-text">{label}</span>
        </label>
      )}
      <input className={cn('input input-bordered w-full', error && 'input-error', className)} ref={ref} {...props} />
      {(error || helperText) && (
        <label className="label">
          <span className={cn('label-text-alt', error && 'text-error')}>{error || helperText}</span>
        </label>
      )}
    </div>
  )
})

Input.displayName = 'Input'
export default Input
