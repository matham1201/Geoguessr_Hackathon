import Link from "next/link";

export default function Footer(){
    return(
        <main className="flex justify-center mt-40 space-x-8 text-gray">
            <Link href="mailto:support@ynov.com">
                Contact
            </Link>
            <div>
                Mentions Légales
            </div>
            <div>
                Cookies
            </div>
        </main>
    )
}