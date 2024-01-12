export async function GET(request: Request) {
    const { searchParams } = new URL(request.url)
    const id = searchParams.get('id')
    var res
    if (id !== null) {
       res = await fetch(`http://localhost:7000/salle/${id}`, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
    } else {
      res = await fetch(`http://localhost:7000/salle`, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
    }
  
    const salle = await res.json()
  
    return Response.json({ salle })
  }