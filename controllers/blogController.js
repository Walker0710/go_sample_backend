const Blog = require('../models/Blog');
const User = require('../models/User');

exports.commentOnBlog = async (req, res) => {
  const { blogId } = req.params;
  const { comment } = req.body;

  try {
    const blog = await Blog.findById(blogId);

    if (!blog) {
      return res.status(404).json({ message: 'Blog not found' });
    }

    const newComment = {
      username: req.user.username,
      text: comment,
      createdAt: new Date(),
    };

    blog.comments.push(newComment);
    await blog.save();

    res.status(201).json({ message: 'Comment added', comment: newComment });
  } catch (err) {
    console.error('Server Error:', err);
    res.status(500).json({ message: 'Server error' });
  }
};


exports.getBlogs = async (req, res) => {
  try {
    // const blogs = await Blog.find().populate('author', 'username');
    const blogs = await Blog.find();
    res.status(200).json(blogs);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.getBlog = async (req, res) => {
  try {
    const blog = await Blog.findById(req.params.id).populate('author', 'username');
    if (!blog) {
      return res.status(404).json({ message: 'Blog not found' });
    }
    res.status(200).json(blog);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.createBlog = async (req, res) => {
  const { title, content } = req.body;

  try {
    const blog = await Blog.create({
      title,
      content,
      author: req.user.username,
    });

    res.status(201).json(blog);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};

exports.commentOnBlog = async (req, res) => {
  const { comment } = req.body;

  try {
    const blog = await Blog.findById(req.params.id);

    if (!blog) {
      return res.status(404).json({ message: 'Blog not found' });
    }

    blog.comments.push({
      username: req.user.username,
      text: comment.text,
    });

    await blog.save();

    res.status(201).json(blog);
  } catch (error) {
    res.status(500).json({ message: 'Server error' });
  }
};
