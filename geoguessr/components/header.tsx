import Link from "next/link";

export default function Header() {
    return(
        <main className="flex justify-between text-gray items-center">
            <img src="/logo_ynovcampus_couleur.png" width={150} className="m-4"/>
            <Link href="/login" className=" flex mr-8 h-fit">
                login
            </Link>
        </main>
    )
}