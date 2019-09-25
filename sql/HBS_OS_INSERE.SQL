create or replace PROCEDURE HBS_OS_INSERE(
            nr_cpf_p      varchar2,
            ds_dano_p             varchar2,
            ds_contato_p varchar2,
            retorno_w OUT varchar2,
            nr_ordem_W OUT number) is

cd_pessoa_fisica_w    varchar2(11);
cd_usuario_w varchar2(20);
cd_setor_usuario_w number;
nr_seq_localizacao_w number;

begin
    select	nvl(max(cd_pessoa_fisica),0)
    into    cd_pessoa_fisica_w
    from	tasy.pessoa_fisica
    where	nr_cpf = NR_CPF_p;
    
    if (cd_pessoa_fisica_w = 0) then 
          retorno_w := 'CPF não cadastrado no TASY, não é possível identificar o responsável';
          RETURN;
    end if;
    
    select nvl(max(u.cd_setor_atendimento),0)
    into cd_setor_usuario_w        
    from usuario u
    where u.cd_pessoa_fisica = cd_pessoa_fisica_w;
    
    select nvl(max(NR_SEQUENCIA) ,0)       
    into nr_seq_localizacao_w
    from MAN_LOCALIZACAO
    where cd_setor = cd_setor_usuario_w;

    select	man_ordem_servico_seq.nextval
    into	nr_ordem_W
    from	dual;
    
    insert	into man_ordem_servico(
        nr_sequencia,
        nr_seq_localizacao,
        nr_seq_equipamento,
        cd_pessoa_solicitante,
        dt_ordem_servico,
        ie_prioridade,
        ie_parado,
        ds_dano_breve,
        dt_atualizacao,
        nm_usuario,
        dt_inicio_desejado,
        dt_conclusao_desejada,
        ds_dano,
        dt_inicio_previsto,
        dt_fim_previsto,
        dt_inicio_real,
        dt_fim_real,
        ie_tipo_ordem,
        ie_status_ordem,
        nr_grupo_planej,
        nr_grupo_trabalho,
        nr_seq_tipo_solucao,
        ds_solucao,
        nm_usuario_exec,
        qt_contador,
        nr_seq_planej,
        nr_seq_tipo_contador,
        nr_seq_estagio,
        cd_projeto,
        nr_seq_etapa_proj,
        dt_reabertura,
        cd_funcao,
        nm_tabela,
        ie_classificacao,
        nr_seq_origem,
        nr_seq_projeto,
        ie_grau_satisfacao,
        nr_seq_indicador,
        nr_seq_causa_dano,
        ie_forma_receb,
        nr_seq_cliente,
        nr_seq_grupo_des,
        nr_seq_grupo_sup,
        nr_seq_superior,
        ie_eficacia,
        dt_prev_eficacia,
        cd_pf_eficacia,
        nr_seq_nao_conform,
        nr_seq_complex,
        dt_atualizacao_nrec,
        nm_usuario_nrec,
        ie_obriga_news,
        nr_seq_meta_pe,
        nr_seq_classif,
        nr_seq_nivel_valor,
        nm_usuario_lib_news,
        dt_libera_news,
        dt_envio_wheb,
        ds_contato_solicitante,
        ie_prioridade_desen,
        ie_prioridade_sup,
        ie_origem_os)
    values(nr_ordem_W,				-- nr_sequencia,
        decode(nr_seq_localizacao_w,0, null, nr_seq_localizacao_w) ,			-- nr_seq_localizacao,
        31,			-- nr_seq_equipamento,
        cd_pessoa_fisica_w,			-- cd_pessoa_solicitante,
        sysdate,		-- dt_ordem_servico,
        'M',				-- ie_prioridade,
        'N',				-- ie_parado,
        substr(ds_dano_p,1,80),		-- ds_dano_breve,
        sysdate,					-- dt_atualizacao,
        'Integrado Site',				-- nm_usuario,
        sysdate,					-- dt_inicio_desejado,
        (sysdate + 15),				-- dt_conclusao_desejada,
        ds_dano_p,				-- ds_dano,
        sysdate,					-- dt_inicio_previsto,
        null,					-- dt_fim_previsto,
        sysdate,					-- dt_inicio_real,
        null,					-- dt_fim_real,
        1,				-- ie_tipo_ordem,
        1,				-- ie_status_ordem,
        41,				-- nr_grupo_planej,
        11,			-- nr_grupo_trabalho,
        null,					-- nr_seq_tipo_solucao,
        null,					-- ds_solucao,
        null, -- nm_usuario_exec,
        null,					-- qt_contador,
        null,					-- nr_seq_planej,
        null,					-- nr_seq_tipo_contador,
        null,				-- nr_seq_estagio,
        null,					-- cd_projeto,
        null,					-- nr_seq_etapa_proj,
        null,					-- dt_reabertura,
        null,					-- cd_funcao,
        null,					-- nm_tabela,
        null,				-- ie_classificacao,
        null,					-- nr_seq_origem,
        null,					-- nr_seq_projeto,
        null,					-- ie_grau_satisfacao,
        null,					-- nr_seq_indicador,
        null,					-- nr_seq_causa_dano,
        null,					-- ie_forma_receb,
        null,					-- nr_seq_cliente,
        null,					-- nr_seq_grupo_des,
        null,					-- nr_seq_grupo_sup,
        null,					-- nr_seq_superior,
        null,					-- ie_eficacia,
        null,					-- dt_prev_eficacia,
        null,					-- cd_pf_eficacia,
        null,					-- nr_seq_nao_conform,
        null,					-- nr_seq_complex,
        null,					-- dt_atualizacao_nrec,
        'Integrado Site',				-- nm_usuario_nrec,
        null,					-- ie_obriga_news,
        null,					-- nr_seq_meta_pe,
        null,					-- nr_seq_classif,
        null,					-- nr_seq_nivel_valor,
        null,					-- nm_usuario_lib_news,
        null,					-- dt_libera_news,
        null,					-- dt_envio_wheb,
        ds_contato_p,					-- ds_contato_solicitante,
        null,					-- ie_prioridade_desen,
        null,					-- ie_prioridade_sup
        '4');					-- ie_origem_os

return;

end HBS_OS_INSERE; 