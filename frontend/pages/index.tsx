import React, { useRef, useState } from 'react'
import SubTest from '@/components/SubTest'
import axios from 'axios'

export default function Home() {
  const [tokens, setToken] = useState('')
  const [userID, setUserID] = useState('')

  const getJWT = async () => {
    const response = await axios.post('http://localhost:8000/getJWT', {
      data: {
        uid: 'asdf',
        pwd: 'aaa',
      },
    })
    const statusOK: boolean = 200 <= response.status && response.status < 300
    console.log(statusOK)
    if (statusOK) {
      setToken((): string => {
        return response.data['token']
      })
    }
  }

  const checkJWT = async () => {
    const config = {
      method: 'get',
      url: 'http://localhost:8000/checkJWT',
      headers: {
        token:
          tokens,
      },
    }
    const response = await axios(config)
    const statusOK: boolean = 200 <= response.status && response.status < 300
    console.log(statusOK)
    if (statusOK) {
      setUserID(() => {
        return response.data
      })
    }
  }

  return (
    <>
      <button onClick={getJWT}>request</button>
      <button onClick={checkJWT}>check JWT</button>
      <p>{userID}</p>
    </>
  )
}
