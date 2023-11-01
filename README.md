# NFT Portal Guide

## Project Layout

### 1. Smart Contracts:

- `contracts/`
- `test/`
- `hardhat.config.ts`
- `package.json`
- `package-lock.json`
- `tsconfig.json`

### 3. Backend API:

- `cmd/`
- `pkg/`
- `go.mod`
- `go.sum`

## Setting Up NFT Metadata

**For local IPFS usage, initiate with:** `$ ipfs daemon`.

1. Add your image to IPFS using: `$ ipfs add file1.jpg`.
2. Update the image URL in `file1.json` to follow this format: `https://ipfs.io/ipfs/`.
3. Upload metadata to IPFS: `$ ipfs add file1.json`.
4. Update the `tokenURIs.json` file.
5. Repeat the above steps for a total of 5 iterations.

If you're not utilizing local IPFS, ensure you omit the IPFS_URL from `docker-compose.yaml`.

## Initiating Hardhat & Contract Deployment

1. Install necessary dependencies: `$ npm install`.
2. Launch a local node using: `$ npx hardhat node`.
3. Deploy your smart contract with: `$ npx hardhat run --network localhost scripts/deploy.ts`.
4. Upon completion, note down the smart contract address.
5. Update the `docker-compose.yaml` file with the CONTRACT_ADDRESS from the previous step.

## Starting the Services

For additional configurations, review this [guide](https://wagmi.sh/react/providers/configuring-chains).

Use the command `$ docker compose up` and then navigate to `http://localhost:3000` in your browser.

## Development Process

1. Set up a local PostgreSQL database and create a `benz` database.
2. Duplicate `.env.sample` to `.env` and adjust the values accordingly.
3. Launch the backend with: `$ godotenv -f .env -- go run ./cmd/api/main.go`.
4. Activate the frontend using: `$ npm run dev`.
