const jwt = require('jsonwebtoken');
const User = require('../models/User');

const auth = async (req, res, next) => {
  try {
    const authHeader = req.header('Authorization');

    if (!authHeader) {
      return res.status(401).json({ message: 'Authorization header is missing' });
    }

    const token = authHeader.replace('Bearer ', '');

    const decoded = jwt.verify(token, process.env.JWT_SECRET);
    const user = await User.findById(decoded.id);

    if (!user) {
      return res.status(401).json({ message: 'Invalid token' });
    }

    req.user = user;
    next();
  } catch (err) {
    console.error('Authentication Error:', err);
    res.status(401).json({ message: 'Authorization failed' });
  }
};

module.exports = auth;
