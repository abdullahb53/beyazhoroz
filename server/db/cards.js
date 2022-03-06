import mongoose from 'mongoose'

const cardsSchema = mongoose.Schema({
    title:{
        type:String,
        required:true
    },
    content:{
        type:String,
        required: true
    },
    img:{
        type:String,
        required: false
    },
    phone:{
        type:String,
        required:true
    },
})

const card = mongoose.model('card', cardsSchema)

export default card