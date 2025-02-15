// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file v1/journal.proto (package v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file v1/journal.proto.
 */
export declare const file_v1_journal: GenFile;

/**
 * @generated from message v1.JournalEntry
 */
export declare type JournalEntry = Message<"v1.JournalEntry"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string codigo = 2;
   */
  codigo: string;

  /**
   * @generated from field: string denominacion = 3;
   */
  denominacion: string;

  /**
   * @generated from field: string empresa_id = 4;
   */
  empresaId: string;

  /**
   * @generated from field: string efe = 5;
   */
  efe: string;

  /**
   * @generated from field: string smvn = 6;
   */
  smvn: string;

  /**
   * @generated from field: string smve = 7;
   */
  smve: string;

  /**
   * @generated from field: string tipo = 8;
   */
  tipo: string;

  /**
   * @generated from field: string tipo_id = 9;
   */
  tipoId: string;

  /**
   * @generated from field: bool esf = 10;
   */
  esf: boolean;

  /**
   * @generated from field: bool erf = 11;
   */
  erf: boolean;

  /**
   * @generated from field: bool ern = 12;
   */
  ern: boolean;

  /**
   * @generated from field: bool efle = 13;
   */
  efle: boolean;

  /**
   * @generated from field: bool ecp = 14;
   */
  ecp: boolean;

  /**
   * @generated from field: string cierre_id = 15;
   */
  cierreId: string;

  /**
   * @generated from field: int32 sequence = 16;
   */
  sequence: number;

  /**
   * @generated from field: string created_at = 17;
   */
  createdAt: string;

  /**
   * @generated from field: string updated_at = 18;
   */
  updatedAt: string;

  /**
   * @generated from field: string deleted_at = 19;
   */
  deletedAt: string;

  /**
   * @generated from field: float iidebe = 20;
   */
  iidebe: number;

  /**
   * @generated from field: float iihaber = 21;
   */
  iihaber: number;

  /**
   * @generated from field: float debe = 22;
   */
  debe: number;

  /**
   * @generated from field: float haber = 23;
   */
  haber: number;
};

/**
 * Describes the message v1.JournalEntry.
 * Use `create(JournalEntrySchema)` to create a new message.
 */
export declare const JournalEntrySchema: GenMessage<JournalEntry>;

/**
 * @generated from message v1.GeneralJournal
 */
export declare type GeneralJournal = Message<"v1.GeneralJournal"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string cuo = 2;
   */
  cuo: string;

  /**
   * @generated from field: string operacion = 3;
   */
  operacion: string;

  /**
   * @generated from field: string descripcion = 4;
   */
  descripcion: string;

  /**
   * @generated from field: string cuenta = 5;
   */
  cuenta: string;

  /**
   * @generated from field: string denominacion = 6;
   */
  denominacion: string;

  /**
   * @generated from field: float debe = 7;
   */
  debe: number;

  /**
   * @generated from field: float haber = 8;
   */
  haber: number;

  /**
   * @generated from field: float suma_18 = 9;
   */
  suma18: number;

  /**
   * @generated from field: float suma_19 = 10;
   */
  suma19: number;

  /**
   * @generated from field: string car = 11;
   */
  car: string;
};

/**
 * Describes the message v1.GeneralJournal.
 * Use `create(GeneralJournalSchema)` to create a new message.
 */
export declare const GeneralJournalSchema: GenMessage<GeneralJournal>;

/**
 * @generated from message v1.RetrieveGeneralJournalRequest
 */
export declare type RetrieveGeneralJournalRequest = Message<"v1.RetrieveGeneralJournalRequest"> & {
  /**
   * @generated from field: string business_id = 1;
   */
  businessId: string;

  /**
   * @generated from field: string period = 2;
   */
  period: string;
};

/**
 * Describes the message v1.RetrieveGeneralJournalRequest.
 * Use `create(RetrieveGeneralJournalRequestSchema)` to create a new message.
 */
export declare const RetrieveGeneralJournalRequestSchema: GenMessage<RetrieveGeneralJournalRequest>;

/**
 * @generated from message v1.RetrieveGeneralJournalResponse
 */
export declare type RetrieveGeneralJournalResponse = Message<"v1.RetrieveGeneralJournalResponse"> & {
  /**
   * @generated from field: repeated v1.GeneralJournal general_journals = 1;
   */
  generalJournals: GeneralJournal[];
};

/**
 * Describes the message v1.RetrieveGeneralJournalResponse.
 * Use `create(RetrieveGeneralJournalResponseSchema)` to create a new message.
 */
export declare const RetrieveGeneralJournalResponseSchema: GenMessage<RetrieveGeneralJournalResponse>;

/**
 * @generated from message v1.RetrieveJournalReportRequest
 */
export declare type RetrieveJournalReportRequest = Message<"v1.RetrieveJournalReportRequest"> & {
  /**
   * @generated from field: string business_id = 1;
   */
  businessId: string;

  /**
   * @generated from field: string period = 2;
   */
  period: string;

  /**
   * @generated from field: bool is_consolidated = 3;
   */
  isConsolidated: boolean;

  /**
   * @generated from field: bool include_close = 4;
   */
  includeClose: boolean;

  /**
   * @generated from field: bool include_cu_ba = 5;
   */
  includeCuBa: boolean;
};

/**
 * Describes the message v1.RetrieveJournalReportRequest.
 * Use `create(RetrieveJournalReportRequestSchema)` to create a new message.
 */
export declare const RetrieveJournalReportRequestSchema: GenMessage<RetrieveJournalReportRequest>;

/**
 * @generated from message v1.RetrieveJournalReportResponse
 */
export declare type RetrieveJournalReportResponse = Message<"v1.RetrieveJournalReportResponse"> & {
  /**
   * @generated from field: repeated v1.JournalEntry journals = 1;
   */
  journals: JournalEntry[];
};

/**
 * Describes the message v1.RetrieveJournalReportResponse.
 * Use `create(RetrieveJournalReportResponseSchema)` to create a new message.
 */
export declare const RetrieveJournalReportResponseSchema: GenMessage<RetrieveJournalReportResponse>;

/**
 * @generated from service v1.JournalService
 */
export declare const JournalService: GenService<{
  /**
   * @generated from rpc v1.JournalService.RetrieveJournalReport
   */
  retrieveJournalReport: {
    methodKind: "unary";
    input: typeof RetrieveJournalReportRequestSchema;
    output: typeof RetrieveJournalReportResponseSchema;
  },
  /**
   * @generated from rpc v1.JournalService.RetrieveGeneralJournal
   */
  retrieveGeneralJournal: {
    methodKind: "unary";
    input: typeof RetrieveGeneralJournalRequestSchema;
    output: typeof RetrieveGeneralJournalResponseSchema;
  },
}>;

