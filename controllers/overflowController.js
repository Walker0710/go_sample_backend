const Overflow = require('../models/Overflow');
const User = require('../models/User');

exports.commentOnOverflow = async (req, res) => {
  const { overflowId } = req.params;
  const { comment } = req.body;

  try {
    const overflow = await Overflow.findById(overflowId);

    if (!overflow) {
      return res.status(404).json({ message: 'overflow not found' });
    }

    const newComment = {
      username: req.user.username,
      text: comment,
      createdAt: new Date(),
    };

    overflow.comments.push(newComment);
    await overflow.save();

    res.status(201).json({ message: 'Comment added', comment: newComment });
  } catch (err) {
    console.error('Server Error:', err);
    res.status(500).json({ message: 'Server error' });
  }
};


exports.getOverflows = async (req, res) => {
  try {
    // const blogs = await Blog.find().populate('author', 'username');
    const overflow = await Overflow.find();
    res.status(200).json(overflow);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.getOverflow = async (req, res) => {
  try {
    const overflow = await Overflow.findById(req.params.id).populate('author', 'username');
    if (!overflow) {
      return res.status(404).json({ message: 'overflow not found' });
    }
    res.status(200).json(overflow);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.createOverflow = async (req, res) => {
  const { title, content } = req.body;

  try {
    const overflow = await Overflow.create({
      title,
      content,
      author: req.user.username,
    });

    res.status(201).json(overflow);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.commentOnOverflow = async (req, res) => {
  const { comment } = req.body;

  try {
    const overflow = await Overflow.findById(req.params.id);

    if (!overflow) {
      return res.status(404).json({ message: 'overflow not found' });
    }

    overflow.comments.push({
      username: req.user.username,
      text: comment.text,
    });

    await overflow.save();

    res.status(201).json(blog);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.getUserOverflows = async (req, res) => {
  try {
    const overflows = await Overflow.find({ author: req.params.username });
    res.json(overflows);
  } catch (error) {
    console.error(error.message);
    res.status(500).send('Server Error');
  }
};
