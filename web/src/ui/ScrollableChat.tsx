import { useRef, useEffect } from 'react'
import { cn } from '@utils/cn'

interface ScrollableChatProps {
  children: React.ReactNode
  className?: string
  autoScroll?: boolean
}

const ScrollableChat = ({ children, className, autoScroll = true }: ScrollableChatProps) => {
  const scrollRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (autoScroll && scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight
    }
  }, [children, autoScroll])

  return (
    <div ref={scrollRef} className={cn('flex flex-col overflow-y-auto', className)}>
      {children}
    </div>
  )
}

export default ScrollableChat
