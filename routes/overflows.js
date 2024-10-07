const express = require('express');
const { getOverflows, getOverflow, createOverflow, commentOnOverflow, getUserOverflows } = require('../controllers/overflowController');
const auth = require('../middlewares/auth');

const router = express.Router();

router.get('/', getOverflows);
router.get('/:id', getOverflow);
router.get('/user/:username', getUserOverflows); // Changed to '/user/:username'
router.post('/', auth, createOverflow);
router.post('/:id/comments', auth, commentOnOverflow);

module.exports = router;
