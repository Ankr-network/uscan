import { TableHeader } from './index';

/**
 * Metadata
 * @class
 */
export class Metadata {
  id: string;
  name: string;
  fileName: string;
  /**
   * Create a Metadata.
   * @param {number} id
   * @param {string} name
   * @param {string} fileName
   */
  constructor(id: string, name: string, fileName: string) {
    this.id = id;
    this.name = name;
    this.fileName = fileName;
  }
}

/**
 * verifyContractMetadata
 * @class
 */
export class VerifyContractMetadata {
  compilerTypes: Metadata[];
  compilerVersions: Metadata[];
  licenseTypes: Metadata[];
  /**
   * Create a Metadata.
   * @param {Metadata[]} compilerTypes
   * @param {Metadata[]} compilerVersions
   * @param {Metadata[]} licenseTypes
   */
  constructor(compilerTypes: Metadata[], compilerVersions: Metadata[], licenseTypes: Metadata[]) {
    this.compilerTypes = compilerTypes;
    this.compilerVersions = compilerVersions;
    this.licenseTypes = licenseTypes;
  }
}

export interface ContractContent {
  id: number;
  contractName: string;
  compilerVersion: string;
  optimization: number;
  runs: number;
  evmVersion: string;
  licenseType: number;
  abi: string;
  metadata: any;
  object: string;
}

export const ContractsHeaderList: TableHeader[] = [
  new TableHeader('Contract Address', 'owner'),
  new TableHeader('Creator', 'creator'),
  new TableHeader('Name', 'name'),
  new TableHeader('Symbol', 'symbol'),
];
