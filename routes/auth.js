const express = require('express');
const { register, login, profile } = require('../controllers/authController');
const auth = require('../middlewares/auth');

const router = express.Router();

router.post('/register', register);
router.post('/login', login);
router.get('/profile', auth, profile);

module.exports = router;
