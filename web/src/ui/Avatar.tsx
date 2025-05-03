import { cn } from '@utils/cn'

interface AvatarProps {
  src?: string
  alt?: string
  size?: 'sm' | 'md' | 'lg'
  online?: boolean
  className?: string
}

const Avatar = ({ src, alt, size = 'md', online = false, className }: AvatarProps) => {
  const sizeClasses = {
    sm: 'w-8 h-8',
    md: 'w-12 h-12',
    lg: 'w-16 h-16',
  }

  return (
    <div className="relative">
      <div className={cn('avatar', online && 'online', className)}>
        <div className={cn('rounded-full', sizeClasses[size])}>
          <img src={src || '/default-avatar.png'} alt={alt || 'User avatar'} />
        </div>
      </div>
    </div>
  )
}

export default Avatar
