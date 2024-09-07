// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file v1/sales.proto (package v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { Pagination, TypeResource } from "./common_pb";

/**
 * Describes the file v1/sales.proto.
 */
export declare const file_v1_sales: GenFile;

/**
 * @generated from message v1.SalesReport
 */
export declare type SalesReport = Message<"v1.SalesReport"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string periodo = 2;
   */
  periodo: string;

  /**
   * @generated from field: string cuo = 3;
   */
  cuo: string;

  /**
   * @generated from field: string identificador_linea = 4;
   */
  identificadorLinea: string;

  /**
   * @generated from field: string fecha_emision = 5;
   */
  fechaEmision: string;

  /**
   * @generated from field: string fec_emision = 6;
   */
  fecEmision: string;

  /**
   * @generated from field: string fecha_vencimiento = 7;
   */
  fechaVencimiento: string;

  /**
   * @generated from field: string fec_venc_pag = 8;
   */
  fecVencPag: string;

  /**
   * @generated from field: string codigo_tipo_cdp = 9;
   */
  codigoTipoCdp: string;

  /**
   * @generated from field: string cod_tipo_cdp = 10;
   */
  codTipoCdp: string;

  /**
   * @generated from field: string serie = 11;
   */
  serie: string;

  /**
   * @generated from field: string num_serie_cdp = 12;
   */
  numSerieCdp: string;

  /**
   * @generated from field: string correlativo = 13;
   */
  correlativo: string;

  /**
   * @generated from field: string num_cdp = 14;
   */
  numCdp: string;

  /**
   * @generated from field: string numero_final = 15;
   */
  numeroFinal: string;

  /**
   * @generated from field: string codigo_tipo_doc_identidad = 16;
   */
  codigoTipoDocIdentidad: string;

  /**
   * @generated from field: string cod_tipo_doc_identidad = 17;
   */
  codTipoDocIdentidad: string;

  /**
   * @generated from field: string num_doc_identidad = 18;
   */
  numDocIdentidad: string;

  /**
   * @generated from field: string num_doc_identidad_client = 19;
   */
  numDocIdentidadClient: string;

  /**
   * @generated from field: string razon_social = 20;
   */
  razonSocial: string;

  /**
   * @generated from field: string nom_razon_social_cliente = 21;
   */
  nomRazonSocialCliente: string;

  /**
   * @generated from field: float exportacion = 22;
   */
  exportacion: number;

  /**
   * @generated from field: float mto_val_fact_expo = 23;
   */
  mtoValFactExpo: number;

  /**
   * @generated from field: float base = 24;
   */
  base: number;

  /**
   * @generated from field: float mto_bi_gravada = 25;
   */
  mtoBiGravada: number;

  /**
   * @generated from field: float desc_base = 26;
   */
  descBase: number;

  /**
   * @generated from field: float mto_dscto_bi = 27;
   */
  mtoDsctoBi: number;

  /**
   * @generated from field: float igv = 28;
   */
  igv: number;

  /**
   * @generated from field: float mto_igv = 29;
   */
  mtoIgv: number;

  /**
   * @generated from field: float desc_igv = 30;
   */
  descIgv: number;

  /**
   * @generated from field: float mto_dscto_igv = 31;
   */
  mtoDsctoIgv: number;

  /**
   * @generated from field: float exonerada = 32;
   */
  exonerada: number;

  /**
   * @generated from field: float mto_exonerado = 33;
   */
  mtoExonerado: number;

  /**
   * @generated from field: float inafecta = 34;
   */
  inafecta: number;

  /**
   * @generated from field: float mto_inafecto = 35;
   */
  mtoInafecto: number;

  /**
   * @generated from field: float isc = 36;
   */
  isc: number;

  /**
   * @generated from field: float mto_isc = 37;
   */
  mtoIsc: number;

  /**
   * @generated from field: float base_ivap = 38;
   */
  baseIvap: number;

  /**
   * @generated from field: float mto_b_i_ivap = 39;
   */
  mtoBIIvap: number;

  /**
   * @generated from field: float ivap = 40;
   */
  ivap: number;

  /**
   * @generated from field: float mto_ivap = 41;
   */
  mtoIvap: number;

  /**
   * @generated from field: float otros = 42;
   */
  otros: number;

  /**
   * @generated from field: float mto_otros_trib = 43;
   */
  mtoOtrosTrib: number;

  /**
   * @generated from field: float total = 44;
   */
  total: number;

  /**
   * @generated from field: float mto_total_cp = 45;
   */
  mtoTotalCp: number;

  /**
   * @generated from field: string codigo_moneda = 46;
   */
  codigoMoneda: string;

  /**
   * @generated from field: string cod_moneda = 47;
   */
  codMoneda: string;

  /**
   * @generated from field: float tipo_cambio = 48;
   */
  tipoCambio: number;

  /**
   * @generated from field: float mto_tipo_cambio = 49;
   */
  mtoTipoCambio: number;

  /**
   * @generated from field: string fecha_cdpm = 50;
   */
  fechaCdpm: string;

  /**
   * @generated from field: string fec_emision_mod = 51;
   */
  fecEmisionMod: string;

  /**
   * @generated from field: string codigo_tipo_cdp_mod = 52;
   */
  codigoTipoCdpMod: string;

  /**
   * @generated from field: string cod_tipo_cdp_mod = 53;
   */
  codTipoCdpMod: string;

  /**
   * @generated from field: string num_serie_cdp_mod = 54;
   */
  numSerieCdpMod: string;

  /**
   * @generated from field: string num_cdp_mod = 55;
   */
  numCdpMod: string;

  /**
   * @generated from field: string numero = 56;
   */
  numero: string;

  /**
   * @generated from field: string num_cdp_mod2 = 57;
   */
  numCdpMod2: string;

  /**
   * @generated from field: string identificador_contrato = 58;
   */
  identificadorContrato: string;

  /**
   * @generated from field: bool error1 = 59;
   */
  error1: boolean;

  /**
   * @generated from field: bool identificador = 60;
   */
  identificador: boolean;

  /**
   * @generated from field: string estado_operacion = 61;
   */
  estadoOperacion: string;

  /**
   * @generated from field: string cod_estado_comprobante = 62;
   */
  codEstadoComprobante: string;

  /**
   * @generated from field: float icbper = 63;
   */
  icbper: number;

  /**
   * @generated from field: float mto_icbp = 64;
   */
  mtoIcbp: number;

  /**
   * @generated from field: string estado_cpe = 65;
   */
  estadoCpe: string;

  /**
   * @generated from field: string observaciones = 66;
   */
  observaciones: string;
};

/**
 * Describes the message v1.SalesReport.
 * Use `create(SalesReportSchema)` to create a new message.
 */
export declare const SalesReportSchema: GenMessage<SalesReport>;

/**
 * @generated from message v1.RetrieveSalesPaginatedReportResponse
 */
export declare type RetrieveSalesPaginatedReportResponse = Message<"v1.RetrieveSalesPaginatedReportResponse"> & {
  /**
   * @generated from field: repeated v1.SalesReport sales = 1;
   */
  sales: SalesReport[];

  /**
   * @generated from field: v1.Pagination pagination = 2;
   */
  pagination?: Pagination;
};

/**
 * Describes the message v1.RetrieveSalesPaginatedReportResponse.
 * Use `create(RetrieveSalesPaginatedReportResponseSchema)` to create a new message.
 */
export declare const RetrieveSalesPaginatedReportResponseSchema: GenMessage<RetrieveSalesPaginatedReportResponse>;

/**
 * @generated from message v1.RetrieveSalesPaginatedReportRequest
 */
export declare type RetrieveSalesPaginatedReportRequest = Message<"v1.RetrieveSalesPaginatedReportRequest"> & {
  /**
   * @generated from field: string business_id = 1;
   */
  businessId: string;

  /**
   * @generated from field: string period = 2;
   */
  period: string;

  /**
   * @generated from field: int32 page = 3;
   */
  page: number;

  /**
   * @generated from field: int32 page_size = 4;
   */
  pageSize: number;
};

/**
 * Describes the message v1.RetrieveSalesPaginatedReportRequest.
 * Use `create(RetrieveSalesPaginatedReportRequestSchema)` to create a new message.
 */
export declare const RetrieveSalesPaginatedReportRequestSchema: GenMessage<RetrieveSalesPaginatedReportRequest>;

/**
 * @generated from message v1.RetrieveSalesResourceReportRequest
 */
export declare type RetrieveSalesResourceReportRequest = Message<"v1.RetrieveSalesResourceReportRequest"> & {
  /**
   * @generated from field: string business_id = 1;
   */
  businessId: string;

  /**
   * @generated from field: string period = 2;
   */
  period: string;

  /**
   * @generated from field: v1.TypeResource type = 3;
   */
  type: TypeResource;
};

/**
 * Describes the message v1.RetrieveSalesResourceReportRequest.
 * Use `create(RetrieveSalesResourceReportRequestSchema)` to create a new message.
 */
export declare const RetrieveSalesResourceReportRequestSchema: GenMessage<RetrieveSalesResourceReportRequest>;

/**
 * @generated from message v1.RetrieveSalesResourceReportResponse
 */
export declare type RetrieveSalesResourceReportResponse = Message<"v1.RetrieveSalesResourceReportResponse"> & {
  /**
   * @generated from field: string url = 1;
   */
  url: string;
};

/**
 * Describes the message v1.RetrieveSalesResourceReportResponse.
 * Use `create(RetrieveSalesResourceReportResponseSchema)` to create a new message.
 */
export declare const RetrieveSalesResourceReportResponseSchema: GenMessage<RetrieveSalesResourceReportResponse>;

/**
 * @generated from service v1.SalesService
 */
export declare const SalesService: GenService<{
  /**
   * @generated from rpc v1.SalesService.RetrieveSalesPaginatedReport
   */
  retrieveSalesPaginatedReport: {
    methodKind: "unary";
    input: typeof RetrieveSalesPaginatedReportRequestSchema;
    output: typeof RetrieveSalesPaginatedReportResponseSchema;
  },
  /**
   * @generated from rpc v1.SalesService.RetrieveSalesResourceReport
   */
  retrieveSalesResourceReport: {
    methodKind: "unary";
    input: typeof RetrieveSalesResourceReportRequestSchema;
    output: typeof RetrieveSalesResourceReportResponseSchema;
  },
}>;

