import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
    const url = "https://dearself.onrender.com/ping";

    // Function to make the request
    const makeRequest = async () => {
        try {
            const response = await fetch(url);
            const data = await response.text();
            console.log(data);
            console.log('ping pong'); // Added console.log to notify when the request is done
        } catch (error) {
            console.error('Error:', error);
        }
    };

    // Initial request
    makeRequest();

    // Schedule automatic requests every 14 minutes
    setInterval(makeRequest, 14 * 60 * 1000);

    res.status(200).json({ message: 'Ping scheduled' });

}
