import { cn } from '@utils/cn'

interface DropdownItem {
  label: string
  onClick: () => void
  icon?: React.ReactNode
}

interface DropdownProps {
  trigger: React.ReactNode
  items: DropdownItem[]
  position?: 'left' | 'right'
  className?: string
}

const Dropdown = ({ trigger, items, position = 'left', className }: DropdownProps) => {
  return (
    <div className={cn('dropdown', position === 'right' && 'dropdown-end', className)}>
      <label tabIndex={0} className="cursor-pointer">
        {trigger}
      </label>
      <ul tabIndex={0} className="dropdown-content z-[1] menu p-2 shadow bg-base-100 rounded-box w-52">
        {items.map((item, index) => (
          <li key={index}>
            <a onClick={item.onClick} className="flex items-center gap-2">
              {item.icon}
              {item.label}
            </a>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default Dropdown
