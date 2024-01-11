"use client"

import Footer from "@/components/footer"
import Header from "@/components/header"
import { useRouter } from "next/navigation"
import React from "react"
import { toast, ToastContainer } from "react-toastify"
import 'react-toastify/dist/ReactToastify.css';

export default function Admin() {

    const [password, setPassword] = React.useState('')
    const router = useRouter()

    function checkPassword(e: { preventDefault: () => void }) {
        e.preventDefault();
        if (password === "log") {
            router.push('/');
        } else {
            toast.error("Wrong password !")
        }
    }

    return (
        <><ToastContainer />
            <main className="grid grid-cols-1 min-h-screen content-between">
                <Header />
                <div className="flex flex-col items-center p-24 ">
                    <form onSubmit={checkPassword}>
                        <input type="password" placeholder="admin's password" required className="p-4 border-4 rounded-lg border-blue focus:border-grayblue focus:outline-none"
                            onChange={e => setPassword(e.target.value)} />
                    </form>
                </div>
                <Footer />
            </main></>
    )
}