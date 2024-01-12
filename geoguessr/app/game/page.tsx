"use client"

import React from "react";
import { useRouter } from "next/navigation";
import { toast, ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';
import Header from "@/components/header";
import Footer from "@/components/footer";

const GeoGuessrPage = () => {
    const [data, setData] = React.useState([]);

    React.useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('/api/score');
                const apiData = await response.json();
                setData(apiData.score);
            } catch (error) {
                console.error('Erreur lors de la récupération des données de l\'API', error);
            }
        };

        fetchData();
    }, []);
    console.log(data)
    const [password, setPassword] = React.useState('')
    const router = useRouter();

    const checkPassword = (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        if (password === "log") {
            router.push('/');
        } else {
            toast.error("Wrong password !");
        }
    }

    return (
        <>
            <ToastContainer />
            <main className="grid grid-cols-1 min-h-screen content-between">
                {data.map((item) => (
                    <li key={item.id}>{item.name}</li>
                ))}
                <Header />
                <div className="flex flex-col items-center p-24">
                    <form onSubmit={checkPassword}>
                        <input
                            type="password"
                            placeholder="admin's password"
                            required
                            className="p-4 border-4 rounded-lg border-blue focus:border-grayblue focus:outline-none"
                            onChange={(e) => setPassword(e.target.value)}
                        />
                        <button
                            type="submit"
                            className="mt-4 p-2 bg-blue text-white rounded-md cursor-pointer"
                        >
                            Enter
                        </button>
                    </form>
                </div>
                <Footer />
            </main>
        </>
    );
};

export default GeoGuessrPage;
