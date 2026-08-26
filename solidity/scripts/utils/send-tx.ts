import { Signer, TransactionReceipt, TransactionResponse } from "ethers";

const nextNonce = new WeakMap<Signer, number>();

export async function sendTx(
  signer: Signer,
  send: (nonce: number) => Promise<TransactionResponse>
): Promise<TransactionReceipt> {
  const provider = signer.provider;
  if (!provider) {
    throw new Error("signer has no provider");
  }

  let next = nextNonce.get(signer);
  if (next === undefined) {
    next = await provider.getTransactionCount(await signer.getAddress(), "latest");
  }

  const tx = await send(next);
  const receipt = await tx.wait();
  if (!receipt) {
    throw new Error(`transaction ${tx.hash} was not mined`);
  }
  nextNonce.set(signer, next + 1);
  return receipt;
}
