import { ButtonHTMLAttributes } from 'react'
import { cn } from '@utils/cn'
import Tooltip from './Tooltip'

interface IconButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'accent' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  tooltip?: string
  tooltipPosition?: 'top' | 'bottom' | 'left' | 'right'
  loading?: boolean
}

const IconButton = ({
  children,
  variant = 'ghost',
  size = 'md',
  tooltip,
  tooltipPosition = 'top',
  loading = false,
  className,
  disabled,
  ...props
}: IconButtonProps) => {
  const sizeClasses = {
    sm: 'btn-sm',
    md: 'btn-md',
    lg: 'btn-lg',
  }

  const button = (
    <button
      className={cn('btn btn-square', `btn-${variant}`, sizeClasses[size], loading && 'loading', className)}
      disabled={disabled || loading}
      {...props}
    >
      {children}
    </button>
  )

  if (tooltip) {
    return (
      <Tooltip content={tooltip} position={tooltipPosition}>
        {button}
      </Tooltip>
    )
  }

  return button
}

export default IconButton
