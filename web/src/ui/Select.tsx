import { SelectHTMLAttributes, forwardRef } from 'react'
import { cn } from '@utils/cn'

interface SelectOption {
  value: string
  label: string
}

interface SelectProps extends Omit<SelectHTMLAttributes<HTMLSelectElement>, 'size'> {
  options: SelectOption[]
  error?: string
  label?: string
  size?: 'sm' | 'md' | 'lg'
}

const Select = forwardRef<HTMLSelectElement, SelectProps>(
  ({ className, error, label, options, size = 'md', ...props }, ref) => {
    const sizeClasses = {
      sm: 'select-sm',
      md: 'select-md',
      lg: 'select-lg',
    }

    return (
      <div className="form-control w-full">
        {label && (
          <label className="label">
            <span className="label-text">{label}</span>
          </label>
        )}
        <select
          className={cn('select select-bordered w-full', sizeClasses[size], error && 'select-error', className)}
          ref={ref}
          {...props}
        >
          {options.map((option) => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
        {error && (
          <label className="label">
            <span className="label-text-alt text-error">{error}</span>
          </label>
        )}
      </div>
    )
  },
)

Select.displayName = 'Select'
export default Select
