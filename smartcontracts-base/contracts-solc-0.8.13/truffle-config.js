module.exports = {
  networks: {
    development: {
      host: process.env.GANACHE_HOSTNAME || '127.0.0.1',
      port: process.env.GANACHE_PORT || 8545,
      network_id: "*",
    },
    secondary: {
      host: process.env.GANACHE_HOSTNAME || '127.0.0.1',
      port: process.env.GANACHE_PORT || 8546,
      network_id: "1451",
    },
  },
  compilers: {
    solc: {
      version: "0.8.13",
    },
  },
};