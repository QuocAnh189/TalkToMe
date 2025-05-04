import { cn } from '@utils/cn'

//assets
import avatar_default_img from '@assets/images/user_default.png'

interface AvatarProps {
  src?: string
  alt?: string
  size?: 'sm' | 'md' | 'lg' | 'xl'
  online?: boolean
  className?: string
}

const Avatar = ({ src, alt, size = 'md', online = false, className }: AvatarProps) => {
  const sizeClasses = {
    sm: 'w-8 h-8',
    md: 'w-12 h-12',
    lg: 'w-16 h-16',
    xl: 'w-32 h-32',
  }

  return (
    <div className="relative">
      <div className={cn('avatar', online && 'online', className)}>
        <div className={cn('rounded-full', sizeClasses[size])}>
          <img src={src || avatar_default_img} alt={alt || 'User avatar'} />
        </div>
      </div>
    </div>
  )
}

export default Avatar
