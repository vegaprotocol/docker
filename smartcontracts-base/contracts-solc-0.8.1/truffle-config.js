module.exports = {
  networks: {
    development: {
      host: process.env.GANACHE_HOSTNAME,
      port: 8545,
      network_id: "*",
    },
  },
  compilers: {
    solc: {
      version: "0.8.1",
    },
  },
};