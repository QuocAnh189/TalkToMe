import { InputHTMLAttributes, forwardRef } from 'react'
import { cn } from '@utils/cn'

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  error?: string
  label?: string
  helperText?: string
  startIcon?: React.ReactNode
  endIcon?: React.ReactNode
  onEndIconClick?: () => void
}

const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, error, label, helperText, startIcon, endIcon, onEndIconClick, ...props }, ref) => {
    return (
      <div className="form-control w-full">
        {label && (
          <label className="label">
            <span className="label-text font-bold">{label}</span>
          </label>
        )}
        <div className="relative bg-amber-200 h-[50px]">
          {startIcon && (
            <div className="absolute left-3 top-1/2 -translate-y-1/2 text-base-content/60 z-10">{startIcon}</div>
          )}
          <input
            className={cn(
              'input input-bordered w-full h-full focus:outline-0',
              startIcon && 'pl-12',
              endIcon && 'pr-12',
              error && 'input-error',
              className,
            )}
            ref={ref}
            {...props}
          />
          {endIcon && (
            <div
              className={cn(
                'z-10 absolute right-3 top-1/2 -translate-y-1/2 text-base-content/60',
                onEndIconClick && 'cursor-pointer hover:text-base-content z-10',
              )}
              onClick={onEndIconClick}
            >
              {endIcon}
            </div>
          )}
        </div>
        {(error || helperText) && (
          <label className="label">
            <span className={cn('label-text-alt', error && 'text-error')}>{error || helperText}</span>
          </label>
        )}
      </div>
    )
  },
)

Input.displayName = 'Input'
export default Input
