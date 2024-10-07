const express = require('express');
const { getBlogs, getBlog, createBlog, commentOnBlog } = require('../controllers/blogController');
const auth = require('../middlewares/auth');

const router = express.Router();

router.get('/', getBlogs);
router.get('/:id', getBlog);
router.post('/', auth, createBlog);
router.post('/:id/comments', auth, commentOnBlog);

module.exports = router;
