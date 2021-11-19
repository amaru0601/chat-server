import { useEffect, useRef, useState } from "react"
import eccrypto from "eccrypto"


//other approach to share ws along the app is a context provider but for small app this is ok
const UseWebsocket = (url) => {
    const [messages, setMessages] = useState([])
    const ws = useRef(null)

    useEffect(() => {
        ws.current = new WebSocket(url)
        ws.current.onmessage = async (e) => {
            const data = JSON.parse(e.data)
            console.log(data.sender)
            console.log(data.body)
            if(data.sender === "Server") {
                setMessages((prev) => [...prev, data.body])
            } else {
                const body = JSON.parse(data.body)
                //console.log(body.pk)
                //console.log(body.message)
                const publicKey = {
                    ciphertext : Buffer(body.message.ciphertext),
                    ephemPublicKey : Buffer(body.message.ephemPublicKey),
                    iv: Buffer(body.message.iv),
                    mac: Buffer(body.message.mac)
                }
                //console.log(publicKey)
                const decrypted = await eccrypto.decrypt(Buffer(body.pk), publicKey)
                //console.log(decrypted)
                setMessages((prev) => [...prev, decrypted.toString()])
            }
            //setMessages((prev) => [...prev, data])
        }

        return () => {
            ws.current.close()
            ws.current = null
        }
    }, [])

    return {
        ws,
        messages
    }
} 

export default UseWebsocket;