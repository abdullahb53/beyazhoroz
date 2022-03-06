import express from 'express'
import mongoose from 'mongoose'
import Card from '../db/cards.js'

const router = express.Router()

//localhost:5000
router.get('/', async (req,res)=> {
    res.json({message: "bu bir get istegi"})
})

router.get('/:id', async (req,res)=> {
    
    try {
        const allCards = await Card.find()
        res.json(allCards)
        
    } catch (error) {
        console.log(error)
    }
})

//gonderi olusturma
router.post('/', async (req,res)=> {

    try {    
        const card = req.body()
        const createdCard = await Card.create(card)
        res.status(201).json(createdCard)
    } catch (error) {
        
    }
    
})

//gonderi olusturma
router.put('/:id', async (req,res)=> {
    res.json({message: "bu bir update istegi"})
})

//gonderi olusturma
router.delete('/:id', async (req,res)=> {
    res.json({message: "bu bir delete istegi"})
})


export default router