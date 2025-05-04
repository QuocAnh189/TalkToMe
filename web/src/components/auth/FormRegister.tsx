//hooks
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '@hooks/useRedux'
import { useSignUpMutation } from '@redux/services/auth'

//ui
import Button from '@ui/Button'
import Input from '@ui/Input'
import Loading from '@ui/Loading'

//components
import toast from 'react-hot-toast'

//icons
import { AiFillEye, AiFillEyeInvisible } from 'react-icons/ai'
import { MdEmail } from 'react-icons/md'
import { FaRegUser } from 'react-icons/fa'
import { BsLockFill } from 'react-icons/bs'

//interfaces
import { SignUpRequest } from '@interfaces/user'

//constants
import { ERole } from '@constants/enum'

//store
import { setAuth } from '@redux/slices/auth.slice'

const initForm: SignUpRequest = {
  email: '',
  name: '',
  password: '',
  avatar: null,
  role: ERole.USER,
}

const FormRegister = () => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()

  const [showPassword, setShowPassword] = useState<boolean>(false)
  const [form, setForm] = useState<SignUpRequest>(initForm)

  const [Register, { isLoading }] = useSignUpMutation()

  const handleChangeForm = (name: string, value: any) => {
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleRegister = async () => {
    const formData = new FormData()
    formData.append('email', form.email)
    formData.append('name', form.name)
    formData.append('password', form.password)
    formData.append('avatar', form.avatar)
    formData.append('role', form.role)

    try {
      const result = await Register(formData).unwrap()

      if (result) {
        dispatch(setAuth(result))
        localStorage.setItem(
          'token',
          JSON.stringify({
            accessToken: result.accessToken,
            refreshToken: result.refreshToken,
          }),
        )
        localStorage.setItem('user', JSON.stringify(result.user))
        toast.success('Register successfully.')
        navigate('/')
      }
    } catch (e: any) {
      toast.error('Something went wrong.')
    }
  }

  return (
    <div className="flex flex-col h-full items-center">
      <div className="mt-4 flex flex-col w-4/5 sm:w-full gap-6">
        <Input
          value={form.email}
          onChange={(e) => handleChangeForm('email', e.target.value)}
          startIcon={<MdEmail />}
          label="Email"
          placeholder="Email"
          type="email"
          className="mb-6"
        />
        <Input
          value={form.name}
          onChange={(e) => handleChangeForm('name', e.target.value)}
          startIcon={<FaRegUser />}
          label="Name"
          placeholder="Name"
          type="text"
          className="mb-6"
        />
        <Input
          value={form.password}
          onChange={(e) => handleChangeForm('password', e.target.value)}
          startIcon={<BsLockFill />}
          label="Password"
          placeholder="Password"
          type={showPassword ? 'text' : 'password'}
          className="mb-6"
          endIcon={showPassword ? <AiFillEye /> : <AiFillEyeInvisible />}
          onEndIconClick={() => setShowPassword(!showPassword)}
        />
        <Input
          label="Avatar"
          placeholder="Avatar"
          type="file"
          className="mb-6"
          onChange={(e: any) => {
            handleChangeForm('avatar', e.target.files[0])
          }}
        />
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
        <Button onClick={handleRegister} className="w-full mt-10" color="#6cb2eb">
          <span className="text-white"> {isLoading ? <Loading /> : 'Sign In'}</span>
        </Button>
      </div>
    </div>
  )
}

export default FormRegister
