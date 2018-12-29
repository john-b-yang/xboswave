#!/bin/sh

RED='\033[0;31m'
GREEN='\033[1;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

ProgName=$(basename $0)
#XBOS_NAMESPACE=
WAVE_DEFAULT_ENTITY=/home/gabe/.wave/gabe.wave.ent

sub_help(){
    echo "Usage: $ProgName <subcommand> [options]\n"
    echo "Subcommands:"
    echo "    mk_namespace  <namespace name>"
    echo "    check_ns_access <namespace name> <entity path?>"
    echo "    mk_driver_ent <driver name> <namespace name>"
    echo ""
    echo "For help with each subcommand run:"
    echo "$ProgName <subcommand> -h|--help"
    echo ""
}

sub_mk_namespace(){
    echo "Making namespace entity"
    # checks if the name argument is there
    if [ -z "$1" ]; then
        echo "Need to provide a name for the entity"
        exit 1;
    fi
	name=$1
	wv mke -e 50y --nopassphrase -o $name.ent
    echo '...skipping passphrase...'
    echo '\n' | wv name --public --attester $WAVE_DEFAULT_ENTITY $name.ent $name
    echo '\n' | wv name --public --attester $name.ent $name.ent $name
    echo '\n' | wv rtprove --subject $WAVE_DEFAULT_ENTITY wavemq:publish,subscribe@$name/*
    if [ $? -ne 0 ]; then
        printf "${YELLOW}No existing access; granting to ${WAVE_DEFAULT_ENTITY} ${NC}\n"
        echo '\n' | wv rtgrant --subject $WAVE_DEFAULT_ENTITY -e 3y --attester $name.ent --indirections 5 wavemq:publish,subscribe@$name/*
    fi
}

sub_check_ns_access(){
    if [ -z "$1" ]; then
        echo "Need to provide a namespace"
        exit 1;
    fi

	ns=$1

    if [ -z "$2" ]; then
        printf "Check access for default ent\n"
        target=$WAVE_DEFAULT_ENTITY
    else
        printf "Check access for ${2}\n"
        target=$2
    fi

    if [ -z "$3" ];then
        resource="*"
    else
        resource=$3
    fi

    echo "Check access to $ns for $target"

    echo '\n' | wv rtprove --subject $target -o proof.pem wavemq:publish,subscribe@$ns/$resource
    if [ $? -ne 0 ]; then
        printf "${RED}-- no access --${NC}\n"
        success=1
    else
        wv verify proof.pem
        if [ $? -eq 0 ]; then
            printf "${GREEN}-- ok --${NC}\n"
            success=0
        else
            printf "${RED}-- no access --${NC}\n"
            success=1
        fi
    fi
}

sub_mk_driver_ent(){
    echo "Making driver entity"
    # checks if the name argument is there
    if [ -z "$1" ]; then
        echo "Need to provide a name for the entity"
        exit 1;
    fi
    if [ -z "$2" ]; then
        echo "Need to provide a name for the namespace"
        exit 1;
    fi

	name=$1
	ns=$2

    if [ ! -f $name.ent ]; then
        printf "${YELLOW}Entity does not exist; creating${NC}\n"
	    wv mke -e 10y --nopassphrase -o $name.ent
        echo '...skipping passphrase...'
        echo '\n' | wv name --public --attester $WAVE_DEFAULT_ENTITY $name.ent $name
        echo '\n' | wv name --public --attester $name.ent $ns.ent $ns
    else
        printf "${YELLOW}Already exists${NC}\n"
    fi

    sub_check_ns_access $ns $name.ent $name/*
    if [ $success -ne 0 ]; then
        printf "${YELLOW}No existing access; granting to ${name} ${NC}\n"
        echo '\n' | wv rtgrant --subject $name -e 3y --attester $WAVE_DEFAULT_ENTITY --indirections 0 wavemq:publish,subscribe@$ns/$name/*
        sub_check_ns_access $ns $name.ent $name/*
    else
        printf "${GREEN}Already has access${NC}\n"
    fi
}

subcommand=$1
case $subcommand in
    "" | "-h" | "--help")
        sub_help
        ;;
    *)
        shift
        sub_${subcommand} $@
        if [ $? = 127 ]; then
            echo "Error: '$subcommand' is not a known subcommand." >&2
            echo "       Run '$ProgName --help' for a list of known subcommands." >&2
            exit 1
        fi
        ;;
esac
