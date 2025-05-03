import { ButtonHTMLAttributes } from 'react'
import { cn } from '@utils/cn'

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'accent' | 'ghost' | 'link'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
  fullWidth?: boolean
  color?: string
}

const Button = ({
  children,
  variant = 'primary',
  size = 'md',
  loading = false,
  fullWidth = false,
  className,
  disabled,
  color,
  ...props
}: ButtonProps) => {
  const sizeClasses = {
    sm: 'btn-sm',
    md: 'btn-md',
    lg: 'btn-lg',
  }

  return (
    <button
      className={cn(
        'btn',
        `btn-${variant}`,
        color && `bg-[${color}]`,
        sizeClasses[size],
        fullWidth && 'w-full',
        loading && 'loading',
        className,
      )}
      disabled={disabled || loading}
      {...props}
    >
      {children}
    </button>
  )
}

export default Button
