import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import {Button} from '@/components/ui/button'
import './App.css'

function App() {
  return (
    <div className="min-h-screen bg-slate-900 flex items-center justify-center">
      <h1 className="text-4xl font-bold text-white">Anchor</h1>
      <Button>Start Focus Session</Button>
    </div>
  )
}

export default App
