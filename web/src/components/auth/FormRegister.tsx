//hooks
import { useState } from 'react'

//ui
import Button from '@ui/Button'
import Input from '@ui/Input'

//icons
import { AiFillEye, AiFillEyeInvisible } from 'react-icons/ai'
import { MdEmail } from 'react-icons/md'
import { FaRegUser } from 'react-icons/fa'
import { BsLockFill } from 'react-icons/bs'
interface Props {}

const FormRegister = (_: Props) => {
  const [showPassword, setShowPassword] = useState<boolean>(false)

  return (
    <div className="flex flex-col h-full items-center">
      <form onSubmit={() => {}} className="mt-4 flex flex-col w-4/5 sm:w-full gap-6">
        <Input startIcon={<MdEmail />} label="Email" placeholder="Email" type="email" className="mb-6" />
        <Input startIcon={<FaRegUser />} label="Name" placeholder="Name" type="text" className="mb-6" />
        <Input
          startIcon={<BsLockFill />}
          label="Password"
          placeholder="Password"
          type={showPassword ? 'text' : 'password'}
          className="mb-6"
          endIcon={showPassword ? <AiFillEye /> : <AiFillEyeInvisible />}
          onEndIconClick={() => setShowPassword(!showPassword)}
        />
        <Input label="Avatar" placeholder="Avatar" type="file" className="mb-6" />
        <div className="flex items-center justify-end">
          <div className="flex items-center gap-1">
            <label htmlFor="remember" className="text-sm space-x-2 text-text">
              You have an account ?
            </label>
            <a href="/login" className="text-sm text-text font-semibold hover:underline">
              Sign In
            </a>
          </div>
        </div>
        <Button type="submit" className="w-full mt-10" color="#6cb2eb">
          <span className="text-white">Sign Up</span>
        </Button>
      </form>
      <div className="w-4/5 sm:w-full mt-3 flex flex-col gap-y-2"></div>
    </div>
  )
}

export default FormRegister
