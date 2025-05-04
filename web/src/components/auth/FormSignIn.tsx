//hooks
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '@hooks/useRedux'
import { useSignInMutation } from '@redux/services/auth'

//ui
import Button from '@ui/Button'
import Input from '@ui/Input'
import Loading from '@ui/Loading'

//components
import toast from 'react-hot-toast'

//icons
import { AiFillEye, AiFillEyeInvisible } from 'react-icons/ai'
import { MdEmail } from 'react-icons/md'
import { BsLockFill } from 'react-icons/bs'

//interfaces
import { SignInRequest } from '@interfaces/user'

//store
import { setAuth } from '@redux/slices/auth.slice'

const initForm: SignInRequest = {
  email: '',
  password: '',
}

const FormLogin = () => {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()
  const [showPassword, setShowPassword] = useState<boolean>(false)

  const [Login, { isLoading }] = useSignInMutation()

  const [form, setForm] = useState<SignInRequest>(initForm)

  const handleChangeForm = (name: string, value: string) => {
    setForm((prev) => ({ ...prev, [name]: value }))
  }

  const handleLogin = async () => {
    try {
      const result = await Login(form).unwrap()
      console.log(result)

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
        toast.success('Login successfully.')
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
          startIcon={<BsLockFill />}
          label="Password"
          placeholder="Password"
          type={showPassword ? 'text' : 'password'}
          className="mb-6"
          value={form.password}
          onChange={(e) => handleChangeForm('password', e.target.value)}
          endIcon={showPassword ? <AiFillEye /> : <AiFillEyeInvisible />}
          onEndIconClick={() => setShowPassword(!showPassword)}
        />
        <div className="flex items-center justify-end">
          <div className="flex items-center gap-1">
            <label htmlFor="remember" className="text-sm space-x-2 text-text">
              You don't have an account ?
            </label>
            <a href="/register" className="text-sm text-text font-semibold hover:underline">
              Sign Up
            </a>
          </div>
        </div>
        <Button onClick={handleLogin} className="w-full mt-10" color="#6cb2eb">
          <span className="text-white">{isLoading ? <Loading /> : 'Sign in'}</span>
        </Button>
      </div>
      <div className="w-4/5 sm:w-full mt-3 flex flex-col gap-y-2"></div>
    </div>
  )
}

export default FormLogin
