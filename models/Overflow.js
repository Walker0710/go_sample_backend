const mongoose = require('mongoose');

const commentSchema = new mongoose.Schema({
  username: { type: String, required: true },
  text: { type: String, required: true },
  // createdAt: { type: Date, default: Date.now },
});

const overflowSchema = new mongoose.Schema({
  title: { type: String, required: true },
  content: { type: String, required: true },
  author: { type: String, required: true },
  createdAt: { type: Date, default: Date.now },
  comments: [commentSchema],
});

const Overflow = mongoose.model('Overflow', overflowSchema, 'Overflow');

module.exports = Overflow;
