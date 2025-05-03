import { cn } from '@utils/cn'

interface TypingIndicatorProps {
  className?: string
}

const TypingIndicator = ({ className }: TypingIndicatorProps) => {
  return (
    <div className={cn('flex items-center gap-1', className)}>
      <div className="w-2 h-2 rounded-full bg-base-content/60 animate-bounce"></div>
      <div className="w-2 h-2 rounded-full bg-base-content/60 animate-bounce [animation-delay:0.2s]"></div>
      <div className="w-2 h-2 rounded-full bg-base-content/60 animate-bounce [animation-delay:0.4s]"></div>
    </div>
  )
}

export default TypingIndicator
