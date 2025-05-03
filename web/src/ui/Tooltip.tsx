import { cn } from '@utils/cn'

interface TooltipProps {
  content: string
  position?: 'top' | 'bottom' | 'left' | 'right'
  children: React.ReactNode
  className?: string
}

const Tooltip = ({ content, position = 'top', children, className }: TooltipProps) => {
  return (
    <div className={cn('tooltip', `tooltip-${position}`, className)} data-tip={content}>
      {children}
    </div>
  )
}

export default Tooltip
