import Footer from "@/components/footer";

export default function Login() {
    return (
        <main className="grid grid-cols-1 min-h-screen content-between">
            <div></div>
            <div className="grid grid-cols-3 content-center ml-80 mr-80">
                <div className="flex flex-col items-center justify-between bg-blue rounded-l-xl pt-16 pb-16 text-white">
                    <img src="/logo_ynov_campus_blanc.svg" width={150}/>
                    <div className="mt-8">Bienvenue dans la communauté Ynov</div>
                    <div className="m-8">Candidats, Etudiants, Professeurs et Entreprises : accédez aux outils et services Ynov Campus à l'aide de votre Compte Ynov !</div>
                    <div>
                      Créer un compte  
                    </div>
                    
                </div>
                <form className="flex flex-col items-center justify-center space-y-4 bg-white rounded-r-xl col-span-2 pt-16 pb-16">
                    <div className="text-blue text-4xl font-bold">Connexion</div>
                    <div className="flex flex-col w-3/4">
                        <label htmlFor="mail">E-mail :</label>
                        <input type="email" id="mail" name="mail" required className="border-2 rounded-lg p-2"/>
                    </div>
                    <div className="flex flex-col w-3/4">
                        <label htmlFor="password">Mot de passe :</label>
                        <input type="password" id="password" name="password" required className="border-2 rounded-lg p-2" />
                    </div>
                    <input type="submit" value="Connexion" className="bg-blue p-2 pr-16 pl-16 rounded-lg hover:bg-white border-4 border-blue hover:text-blue text-white" />
                </form>
            </div>
            <Footer />
        </main>
    )
}